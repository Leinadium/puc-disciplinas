import { writable } from "svelte/store";
import type { Writable } from "svelte/store";

/** o nome do usuario, se estiver logado, ou nulo */
export const userStore: Writable<string | null> = writable(null);
export const hasCurriculo: Writable<boolean> = writable(false);