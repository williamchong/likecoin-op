import en from './i18n/en.json';
import zhHantHK from './i18n/zh-Hant-HK.json';

export default {
  legacy: false,
  locales: ['en', 'zh-Hant-HK'],
  defaultLocale: 'zh-Hant-HK',
  vueI18n: {
    fallbackLocale: 'zh-Hant-HK',
    messages: {
      en,
      'zh-Hant-HK': zhHantHK,
    },
  },
};
