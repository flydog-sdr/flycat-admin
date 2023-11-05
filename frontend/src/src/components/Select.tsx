import { ChangeEvent, Component } from "react";

export interface OptionItem {
    name: string;
    value: string;
    selected?: boolean;
}

export interface SelectProps {
    readonly id?: string;
    readonly label: string;
    readonly options: OptionItem[];
    readonly onChange?: (event: ChangeEvent<HTMLSelectElement>) => void;
}

export default class Select extends Component<SelectProps> {
    render() {
        const id = `input-${Math.random().toString(36).substring(2, 15)}`;
        const { label, options, onChange } = this.props;

        return (
            <div className="w-full m-2 p-3 border-solid rounded-md border-2 flex flex-wrap hover:bg-gray-100">
                <label
                    className="block mb-2 text-sm font-medium text-gray-900"
                    htmlFor={id}
                >
                    {label}
                </label>
                <select
                    id={id}
                    onChange={onChange}
                    defaultValue={"_defaultValue_"}
                    className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg block w-full p-2"
                >
                    <option value={"_defaultValue_"} disabled>
                        Choose a value
                    </option>
                    {options.map((item: OptionItem, index: number) => (
                        <option key={index} value={item.value}>
                            {item.name}
                        </option>
                    ))}
                </select>
            </div>
        );
    }
}
