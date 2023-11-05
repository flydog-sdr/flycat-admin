import { OptionItem } from "../components/Select";

const generateAntenna = ({
    start,
    end,
}: {
    start: number;
    end: number;
}): OptionItem[] => {
    const array: OptionItem[] = [];

    for (let i = start-1; i < end; i++) {
        array.push({
            name: `RF-IN${i}`,
            value: i.toString(),
        });
    }

    return array;
};

export default generateAntenna;
