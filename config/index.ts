// [框架] Nuxt Config，不要删除或修改
import type { NuxtConfig } from "nuxt/config";
import misc from "./[Framework]misc";
import nitro from "./[Framework]nitro";
import vite from "./[Framework]vite";

export default { nitro, vite, ...misc } as NuxtConfig;
