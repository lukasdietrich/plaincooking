import { init, register, getLocaleFromNavigator } from 'svelte-i18n';

const locales = ['en', 'de'];
const defaultLocale = 'en';

locales.forEach((locale) => {
	register(locale, () => import(`./locales/${locale}.json`));
});

init({
	fallbackLocale: defaultLocale,
	initialLocale: getLocaleFromNavigator()
});

export { t, waitLocale } from 'svelte-i18n';
