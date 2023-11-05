import { Component, ReactNode } from "react";

export interface PanelProps {
    title: string;
    children: ReactNode;
}

export default class Panel extends Component<PanelProps> {
    render() {
        const { children, title } = this.props;

        return (
            <div className="w-full p-3">
                <div className="bg-white border rounded shadow">
                    <div className="border-b p-3">
                        <h5 className="font-bold uppercase text-gray-600">
                            {title}
                        </h5>
                    </div>
                    <div className="p-5 flex flex-col">{children}</div>
                </div>
            </div>
        );
    }
}
