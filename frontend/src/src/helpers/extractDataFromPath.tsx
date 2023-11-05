const extractDataFromPath = (obj: any, path: string): any => {
    const keys = path.split("->");

    let value = obj;
    for (const key of keys) {
        if (value.hasOwnProperty(key)) {
            value = value[key];
        } else {
            return undefined;
        }
    }

    return value;
};

export default extractDataFromPath;
