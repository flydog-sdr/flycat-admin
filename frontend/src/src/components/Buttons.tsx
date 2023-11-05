import { Component } from "react";
import Button, { ButtonProps } from "./Button";

export interface ButtonsItem extends ButtonProps {
    readonly floating: "left" | "right";
}

export interface ButtonsProps {
    readonly items: ButtonsItem[];
}

export default class Buttons extends Component<ButtonsProps> {
    render() {
        return (
            <div className="flex p-2 flex-wrap">
                {this.props.items.map(
                    (item: ButtonsItem, index: number) =>
                        item.floating === "left" && (
                            <Button
                                key={index}
                                icon={item.icon}
                                label={item.label}
                                onClick={item.onClick}
                                textColor={item.textColor}
                                hoverColor={item.hoverColor}
                                borderColor={item.borderColor}
                                backgroundColor={item.backgroundColor}
                            />
                        )
                )}

                <div className="flex-auto flex flex-row-reverse flex-wrap">
                    {this.props.items.map(
                        (item: ButtonsItem, index: number) =>
                            item.floating === "right" && (
                                <Button
                                    key={index}
                                    icon={item.icon}
                                    label={item.label}
                                    onClick={item.onClick}
                                    textColor={item.textColor}
                                    hoverColor={item.hoverColor}
                                    borderColor={item.borderColor}
                                    backgroundColor={item.backgroundColor}
                                />
                            )
                    )}
                </div>
            </div>
        );
    }
}
