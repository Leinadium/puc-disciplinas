import { writable } from "svelte/store";
import type { Writable, Readable } from "svelte/store";

/** o nome do usuario, se estiver logado, ou nulo */
export const userStore: Writable<string | null> = writable(null);
export const hasCurriculo: Writable<boolean> = writable(false);

const events: any = (store: any) => ({
    ...store,
    subscribe(subscriber: any) {
        let initial = true;
        const unsubscribe = store.subscribe(($x: any) => {
            if (!initial) {
                subscriber($x);
            }
        })
        initial = false;
        return unsubscribe;
    }
})

export const userEvent: Readable<string | null> = events(userStore);