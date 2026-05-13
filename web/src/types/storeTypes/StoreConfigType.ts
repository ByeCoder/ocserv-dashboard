export interface ConfigState {
    setup: boolean;
    googleCaptchaSiteKey: string;
}

export interface Release {
    Current: string;
    Latest: string;
}

export interface ServerState {
    OcservVersion: string;
    OcctlVersion: string;
    Status: string;
    Release: Release;
}
