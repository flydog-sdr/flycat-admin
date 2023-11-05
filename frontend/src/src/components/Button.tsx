import { Component, MouseEventHandler } from "react";

export interface ButtonProps {
    readonly icon: string;
    readonly label: string;
    readonly textColor: string;
    readonly hoverColor: string;
    readonly borderColor: string;
    readonly backgroundColor: string;
    readonly onClick: MouseEventHandler<HTMLButtonElement>;
}

export default class Button extends Component<ButtonProps> {
    render() {
        const {
            icon,
            label,
            textColor,
            hoverColor,
            borderColor,
            backgroundColor,
            onClick,
        } = this.props;

        return (
            <button
                className={`text-base m-1 hover:scale-110 focus:outline-none flex justify-center px-4 py-2 rounded font-bold ${hoverColor} ${backgroundColor} ${textColor} ${borderColor} border duration-200`}
                onClick={onClick}
            >
                <div className="flex flex-row gap-2">
                    <img src={icon} alt="" width={20} />
                    {label}
                </div>
            </button>
        );
    }
}
