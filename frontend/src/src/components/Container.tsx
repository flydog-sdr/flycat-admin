import { Component, ReactNode } from "react";

export interface ContainerProps {
    children: ReactNode;
}

class Container extends Component<ContainerProps> {
    render() {
        const { children } = this.props;

        return (
            <div className="container w-full mx-auto pt-20">
                <div className="w-full px-4 md:px-8 md:mt-8 mb-16 text-gray-800 leading-normal">
                    {children}
                </div>
            </div>
        );
    }
}

class ItemContainer extends Component<ContainerProps> {
    render() {
        const { children } = this.props;

        return (
            <div className="flex flex-row flex-wrap flex-grow mt-2">
                {children}
            </div>
        );
    }
}

export default Container;
export { ItemContainer };
