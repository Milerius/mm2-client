// Use ES module import syntax to import functionality from the module
// that we have compiled.
//
// Note that the `default` import is an initialization function which
// will "boot" the module and make it ready to use. Currently browsers
// don't support natively imported WebAssembly as an ES module, but
// eventually the manual initialization won't be required!
import init, { mm2_main, mm2_main_status, mm2_rpc, LogLevel, Mm2MainErr, MainStatus, Mm2RpcErr } from "./deps/pkg/mm2.js";

const LOG_LEVEL = LogLevel.Info;

// Loads the wasm file, so we use the
// default export to inform it where the wasm file is located on the
// server, and then we wait on the returned promise to wait for the
// wasm to be loaded.
window.init_wasm = async function() {
    try {
        await init();
    } catch (e) {
        alert(`Oops: ${e}`);
    }
}

window.run_mm2 = async function(params) {
    let config = {
        conf: JSON.parse(params),
        log_level: LOG_LEVEL,
    }

    // run an MM2 instance
    try {
        mm2_main(config, handle_log);
    } catch (e) {
        switch (e) {
            case Mm2MainErr.AlreadyRuns:
                alert("MM2 already runs, please wait...");
                return;
            case Mm2MainErr.InvalidParams:
                alert("Invalid config");
                return;
            case Mm2MainErr.NoCoinsInConf:
                alert("No 'coins' field in config");
                return;
            default:
                alert(`Oops: ${e}`);
                return;
        }
    }
}

window.rpc_request = async function(request_js) {
    try {
        let reqJson = JSON.parse(request_js);
        const response = await mm2_rpc(reqJson);
        return JSON.stringify(response);
    } catch (e) {
        switch (e) {
            case Mm2RpcErr.NotRunning:
                alert("MM2 is not running yet");
                break;
            case Mm2RpcErr.InvalidPayload:
                alert(`Invalid payload: ${request_js}`);
                break;
            case Mm2RpcErr.InternalError:
                alert(`An MM2 internal error`);
                break;
            default:
                alert(`Unexpected error: ${e}`);
                break;
        }

        return e;
    }
}

function handle_log(level, line) {
    switch (level) {
        case LogLevel.Off:
            break;
        case LogLevel.Error:
            console.error(line);
            break;
        case LogLevel.Warn:
            console.warn(line);
            break;
        case LogLevel.Info:
            console.info(line);
            break;
        case LogLevel.Debug:
            console.log(line);
            break;
        case LogLevel.Trace:
        default:
            // The console.trace method outputs some extra trace from the generated JS glue code which we don't want.
            console.debug(line);
            break;
    }
}

function mm2_status() {
    return mm2_main_status();
}