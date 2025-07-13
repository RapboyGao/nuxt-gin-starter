// [框架] Nuxt Config，不要删除或修改
import type { NuxtConfig } from "nuxt/config";
import misc from "./misc";
import nitro from "./nitro";
import vite from "./vite";

export default { nitro, vite, ...misc } as NuxtConfig;
