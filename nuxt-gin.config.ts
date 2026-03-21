import createNuxtGinConfig from 'nuxt-gin-tools/src/nuxt-gin';

export default createNuxtGinConfig({
  dev: {
    killPortBeforeDevelop: true,
    cleanupBeforeDevelop: false,
  },
  pack: {},
});
