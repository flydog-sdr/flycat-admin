import { Component } from "react";
import Navbar from "./components/Navbar";
import Card, { CardProps } from "./components/Card";
import Container, { ItemContainer } from "./components/Container";
import Divider from "./components/Divider";
import { errorAlert, successAlert, toastAlert } from "./helpers/swalPopups";
import Panel from "./components/Panel";
import AppConfig from "./config";
import Select, { SelectProps } from "./components/Select";
import generateGain from "./helpers/generateGain";
import Buttons, { ButtonsProps } from "./components/Buttons";
import generateAntenna from "./helpers/generateAntenna";
import Polling from "./components/Polling";
import axios, { AxiosResponse } from "axios";
import extractDataFromPath from "./helpers/extractDataFromPath";
import getApiFullPath from "./helpers/getApiFullPath";

interface State {
    cards: CardProps[];
    selections: SelectProps[];
    buttons: ButtonsProps;
    userActions: {
        gain: number | null;
        filter: number | null;
        antenna: number | null;
    };
}

export default class App extends Component<{}, State> {
    constructor(props: object) {
        super(props);
        this.state = {
            userActions: {
                gain: null,
                filter: null,
                antenna: null,
            },
            cards: [
                {
                    id: "cpu->percent",
                    title: "CPU Usage",
                    value: "0.00",
                    unit: "%",
                    icon: require("./assets/microchip-solid.svg").default,
                    backgroundColor: "bg-amber-600",
                },
                {
                    id: "memory->percent",
                    title: "Memory Usage",
                    value: "0.00",
                    unit: "%",
                    icon: require("./assets/memory-solid.svg").default,
                    backgroundColor: "bg-pink-600",
                },
                {
                    id: "controller->gain",
                    title: "Receiver Gain",
                    value: "0",
                    unit: "dB",
                    icon: require("./assets/arrow-up-right-dots-solid.svg")
                        .default,
                    backgroundColor: "bg-sky-600",
                },
                {
                    id: "controller->filter",
                    title: "Notch Filter",
                    value: "OFF",
                    icon: require("./assets/square-caret-right-solid.svg")
                        .default,
                    backgroundColor: "bg-lime-600",
                },
                {
                    id: "controller->antenna",
                    title: "Antenna NO.",
                    value: "0",
                    unit: "#",
                    icon: require("./assets/tower-broadcast-solid.svg").default,
                    backgroundColor: "bg-indigo-600",
                },
                {
                    id: "uptime",
                    title: "Server Uptime",
                    value: "0",
                    unit: "days",
                    icon: require("./assets/clock-solid.svg").default,
                    backgroundColor: "bg-teal-600",
                },
            ],
            selections: [
                {
                    id: "gain",
                    label: "Receiver Gain",
                    options: generateGain({
                        start: -11,
                        end: 34,
                    }),
                    onChange: (event) => {
                        this.handleOnChange("gain", event.currentTarget.value);
                    },
                },
                {
                    id: "filter",
                    label: "Notch Filter",
                    options: [
                        {
                            name: "OFF",
                            value: "0",
                        },
                        {
                            name: "ON",
                            value: "1",
                        },
                    ],
                    onChange: (event) => {
                        this.handleOnChange(
                            "filter",
                            event.currentTarget.value
                        );
                    },
                },
                {
                    id: "antenna",
                    label: "Receiver Antenna",
                    options: generateAntenna({
                        start: 1,
                        end: 3,
                    }),
                    onChange: (event) => {
                        this.handleOnChange(
                            "antenna",
                            event.currentTarget.value
                        );
                    },
                },
            ],
            buttons: {
                items: [
                    {
                        floating: "right",
                        label: "Save & Apply",
                        textColor: "text-white",
                        backgroundColor: "bg-green-600",
                        hoverColor: "hover:bg-green-700",
                        borderColor: "border-green-600",
                        onClick: this.handleSaveApply,
                        icon: require("./assets/gears-solid.svg").default,
                    },
                ],
            },
        };
    }

    handleOnChange = (id: string, value: string): void => {
        const { userActions } = this.state;

        switch (id) {
            case "gain":
                userActions.gain = parseInt(value);
                break;
            case "filter":
                userActions.filter = parseInt(value);
                break;
            case "antenna":
                userActions.antenna = parseInt(value);
                break;
        }

        this.setState({
            userActions,
        });
    };

    handleSaveApply = async (): Promise<void> => {
        const { userActions } = this.state;
        for (let key in userActions) {
            switch (key) {
                case "gain":
                    if (userActions.gain === null) {
                        errorAlert({
                            title: "Error",
                            html: "Gain is not set yet.",
                        });
                        return;
                    }
                    break;
                case "filter":
                    if (userActions.filter === null) {
                        errorAlert({
                            title: "Error",
                            html: "Notch filter is not set yet.",
                        });
                        return;
                    }
                    break;
                case "antenna":
                    if (userActions.antenna === null) {
                        errorAlert({
                            title: "Error",
                            html: "Antenna is not set yet.",
                        });
                        return;
                    }
                    break;
            }
        }

        try {
            await axios(
                getApiFullPath(AppConfig.api_settings.production, "controller"),
                { method: "POST", data: userActions }
            );
            successAlert({
                title: "Success",
                html: "Successfully saved and applied.",
            });
        } catch {
            errorAlert({
                title: "Error",
                html: "Something went wrong.",
            });
        }
    };

    onFetch = (url: string): Promise<AxiosResponse<any, any>> =>
        axios(url, { method: "GET" });

    onSuccess = (res: any): any => {
        const {
            data: { data },
        } = res;
        this.setState({
            cards: this.state.cards.map((item: CardProps) => {
                const { id } = item;
                let value = extractDataFromPath(data, id);

                switch (id) {
                    case "cpu->percent":
                    case "memory->percent":
                        value = value.toFixed(2);
                        break;
                    case "controller->filter":
                        value = value === 0 ? "OFF" : "ON";
                        break;
                    case "uptime":
                        value = Math.floor(value / 86400);
                        break;
                }

                return {
                    ...item,
                    value,
                };
            }),
        });

        return res;
    };

    onError = () => {
        toastAlert({
            title: "Error",
            html: "Error while fetching data from server.",
            icon: "error",
            timer: 3000,
        });
    };

    componentDidMount(): void {
        const { title } = AppConfig.site_settings;
        document.title = title;
    }

    render() {
        const { api_settings, site_settings } = AppConfig;

        return (
            <>
                <Navbar title={site_settings.title} />

                <Container>
                    <ItemContainer>
                        <Polling
                            interval={3000}
                            onFetch={this.onFetch}
                            onError={this.onError}
                            onSuccess={this.onSuccess}
                            url={getApiFullPath(
                                api_settings.production,
                                "status"
                            )}
                        >
                            {this.state.cards.map(
                                (item: CardProps, index: number) => (
                                    <Card
                                        key={index}
                                        id={item.id}
                                        icon={item.icon}
                                        title={item.title}
                                        value={item.value}
                                        unit={item.unit}
                                        backgroundColor={item.backgroundColor}
                                    />
                                )
                            )}
                        </Polling>
                    </ItemContainer>

                    <Divider />

                    <ItemContainer>
                        <Panel title="FlyCat Controller">
                            {this.state.selections.map(
                                (item: SelectProps, index: number) => (
                                    <Select
                                        key={index}
                                        label={item.label}
                                        options={item.options}
                                        onChange={item.onChange}
                                    />
                                )
                            )}
                            <Buttons items={this.state.buttons.items} />
                        </Panel>
                    </ItemContainer>
                </Container>
            </>
        );
    }
}
