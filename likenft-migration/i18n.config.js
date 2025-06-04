import en from './i18n/en.json';
import { DEFAULT_LOCALE, Locales } from './i18n/locale';
import zhHantHK from './i18n/zh-Hant-HK.json';

export default {
  legacy: false,
  locales: Locales,
  defaultLocale: DEFAULT_LOCALE,
  vueI18n: {
    fallbackLocale: DEFAULT_LOCALE,
    messages: {
      en,
      'zh-Hant-HK': zhHantHK,
    },
  },
};
