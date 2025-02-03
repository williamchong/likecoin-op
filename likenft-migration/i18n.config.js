import en from './i18n/en.json';
import zhHantHK from './i18n/zh-Hant-HK.json';

export default {
  legacy: false,
  locales: ['en', 'zh-Hant-HK'],
  defaultLocale: 'en',
  vueI18n: {
    fallbackLocale: 'en',
    messages: {
      en,
      'zh-Hant-HK': zhHantHK,
    },
  },
};
