import "./monaco";
import App from './views/App.svelte';


const app = new App({
    target: document.body,
});

// TODO: viqueen - figure out how to handle this
// Wails.Init(() => {
//     Wails.Events.Once("wombat:init", ({build_mode}) => {
//         if (build_mode === "prod") window.addEventListener('contextmenu', e => e.preventDefault());
//     });
//
//     // TODO: v2
//     // Hack until we have wails v2 and can use CSS variables again
//     window.isWin = window.navigator.platform.startsWith("Win")
//
//     app = new App({
//         target: document.body,
//     });
// });

export default app;
