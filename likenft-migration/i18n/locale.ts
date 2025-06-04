import { enUS, Locale as DateFnsLocale, zhHK } from 'date-fns/locale';

export type Locale = 'en' | 'zh-Hant-HK';

export const Locales: Locale[] = ['en', 'zh-Hant-HK'];

export const DEFAULT_LOCALE: Locale = 'zh-Hant-HK';

export function getDateFnsLocales(locale: string): DateFnsLocale {
  switch (locale as Locale) {
    case 'en':
      return enUS;
    case 'zh-Hant-HK':
      return zhHK;
  }
  return enUS;
}

export function getDateFmt(locale: string): string {
  switch (locale as Locale) {
    case 'en':
      return 'dd MMM, yyyy HH:mm:ss';
    case 'zh-Hant-HK':
      return 'yyyy/MM/dd HH:mm:ss';
  }
  return 'dd MMM, yyyy HH:mm:ss';
}
