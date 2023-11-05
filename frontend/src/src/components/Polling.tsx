import { Component } from "react";
import ReactPolling from "react-polling";

interface Props {
    readonly url: string;
    readonly interval: number;
    readonly onError: () => void;
    readonly onFetch: (url: string) => void;
    readonly onSuccess: (res: any) => void;
    readonly children: React.ReactNode;
}

export default class Polling extends Component<Props> {
    render() {
        const { url, interval, children, onFetch, onSuccess, onError } =
            this.props;

        return (
            <ReactPolling
                url={url}
                method={"GET"}
                retryCount={2}
                interval={interval}
                promise={onFetch}
                onFailure={onError}
                onSuccess={onSuccess}
                render={() => children}
            />
        );
    }
}
