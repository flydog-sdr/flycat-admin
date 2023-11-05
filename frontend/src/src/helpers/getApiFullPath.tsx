import AppConfig from "../config";

const getApiFullPath = (production: boolean, api: string) => {
    const baseUrl = production
        ? `${window.location.protocol}//${window.location.host}`
        : AppConfig.api_settings.development;
    const apiPath = `/api/${AppConfig.api_settings.version}/${api}`;
    return `${baseUrl}${apiPath}`;
};

export default getApiFullPath;
