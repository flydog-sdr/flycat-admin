import { Component } from "react";

export interface CardProps {
    value: string;
    readonly id: string;
    readonly icon: string;
    readonly title: string;
    readonly unit?: string;
    readonly backgroundColor: string;
}

export default class Card extends Component<CardProps> {
    render() {
        const { value, icon, title, unit, backgroundColor } = this.props;

        return (
            <div className="w-full md:w-1/2 p-3">
                <div className="bg-white border rounded shadow p-2">
                    <div className="flex flex-row items-center">
                        <div className="flex-shrink pr-4">
                            <div className={`rounded p-3 ${backgroundColor}`}>
                                <img
                                    width={30}
                                    height={30}
                                    src={icon}
                                    alt="icon"
                                />
                            </div>
                        </div>
                        <div className="flex-1 text-right md:text-center">
                            <h5 className="font-bold uppercase text-gray-500">
                                {title}
                            </h5>
                            <h3 className="font-bold text-3xl">
                                {value} {unit}
                            </h3>
                        </div>
                    </div>
                </div>
            </div>
        );
    }
}
