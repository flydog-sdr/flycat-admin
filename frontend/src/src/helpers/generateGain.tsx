import { OptionItem } from "../components/Select";

const generateGain = ({
    start,
    end,
}: {
    start: number;
    end: number;
}): OptionItem[] => {
    const array: OptionItem[] = [];

    for (let i = start; i <= end; i++) {
        array.push({
            name: `${i} dB`,
            value: i.toString(),
            selected: false,
        });
    }

    return array;
};

export default generateGain;