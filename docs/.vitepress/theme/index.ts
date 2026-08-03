import type { Theme } from 'vitepress';
import DefaultTheme from 'vitepress/theme';
import "@catppuccin/vitepress/theme/frappe/blue.css";

import './custom.css';

export default {
	extends: DefaultTheme,
} satisfies Theme;
