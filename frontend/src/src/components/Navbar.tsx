import { Component } from "react";
import QuestionIcon from "../assets/circle-question-solid.svg";

export interface NavbarProps {
    readonly title: string;
    readonly help?: () => void;
}

export default class Navbar extends Component<NavbarProps> {
    render() {
        const { title, help } = this.props;

        return (
            <nav className="bg-white fixed w-full z-10 shadow p-4">
                <div className="w-full container mx-auto flex flex-wrap items-center pt-2 pb-3 md:pb-0">
                    <div className="w-1/2 pl-2 md:pl-5">
                        <a
                            className="text-gray-900 text-base xl:text-xl no-underline hover:no-underline font-bold"
                            href="."
                        >
                            {title}
                        </a>
                    </div>
                    <div className="w-1/2 pr-2">
                        <div className="flex relative float-right">
                            <div className="relative text-sm">
                                <button onClick={help}>
                                    <img
                                        className="inline-block m-1"
                                        alt="?"
                                        width={15}
                                        src={QuestionIcon}
                                    />
                                    <span className="hidden md:inline-block">
                                        Help
                                    </span>
                                </button>
                            </div>
                        </div>
                    </div>
                </div>
            </nav>
        );
    }
}
