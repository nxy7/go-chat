export function WebsocketRelativeAddr(path: string) {
    var loc = window.location, ws_uri;
    if (loc.protocol === "https:") {
        ws_uri = "wss://";
    } else {
        ws_uri = "ws://";
    }
    ws_uri += "//" + loc.host + path;
    return ws_uri
}