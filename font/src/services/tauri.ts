declare global {
  interface Window {
    __TAURI__: any;
  }
}

export function isInTauri(): boolean {
  return typeof window.__TAURI__!== 'undefined'
}
