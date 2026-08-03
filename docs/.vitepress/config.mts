import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: 'gotsrpc',
  description: 'CLI utility to generate go and typescript RPC calls easily',
	lang: "en-US",
	cleanUrls: true,
	lastUpdated: true,
	appearance: "dark",
	ignoreDeadLinks: false,
  base: '/gotsrpc/',
	sitemap: {
		hostname: 'https://foomo.github.io/gotsrpc',
	},
  themeConfig: {
		// https://vitepress.dev/reference/default-theme-config
		logo: '/logo.png',
		outline: [2, 4],
    nav: [
      { text: 'Guide', link: '/guide/getting-started' },
      { text: 'Reference', link: '/reference/cli/gotsrpc' },
    ],
    sidebar: [
      {
        text: 'Guide',
        items: [
          { text: 'Overview', link: '/' },
          { text: 'Getting Started', link: '/guide/getting-started' },
          { text: 'Writing Services', link: '/guide/writing-services' },
          { text: 'Configuration', link: '/guide/configuration' },
        ],
      },
      {
        text: 'Reference',
        items: [
          { text: 'CLI: gotsrpc', link: '/reference/cli/gotsrpc' },
          { text: 'gotsrpc generate', link: '/reference/cli/gotsrpc_generate' },
          { text: 'gotsrpc version', link: '/reference/cli/gotsrpc_version' },
        ],
      },
			{
				text: 'Contributing',
				collapsed: true,
				items: [
					{
						text: "Guideline",
						link: '/CONTRIBUTING.md',
					},
					{
						text: "Code of conduct",
						link: '/CODE_OF_CONDUCT.md',
					},
					{
						text: "Security guidelines",
						link: '/SECURITY.md',
					},
				],
			},
		],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/foomo/gotsrpc' },
    ],
		editLink: {
			pattern: 'https://github.com/foomo/gotsrpc/edit/main/docs/:path',
		},
		search: {
			provider: 'local',
		},
		footer: {
			message: 'Made with ♥ <a href="https://www.foomo.org">foomo</a> by <a href="https://www.bestbytes.com">bestbytes</a>',
		},
  },
	markdown: {
		// https://github.com/vuejs/vitepress/discussions/3724
		theme: {
			light: 'catppuccin-latte',
			dark: 'catppuccin-frappe',
		}
	},
	head: [
		['meta', { name: 'theme-color', content: '#ffffff' }],
		['link', { rel: 'icon', href: '/logo.png' }],
		['meta', { name: 'author', content: 'foomo by bestbytes' }],
		// OpenGraph
		['meta', { property: 'og:title', content: 'foomo/gotsrpc' }],
		[
			'meta',
			{
				property: 'og:image',
				content: 'https://github.com/foomo/gotsrpc/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta',
			{
				property: 'og:description',
				content: 'CLI utility to generate go and typescript RPC calls easily',
			},
		],
		['meta', { name: 'twitter:card', content: 'summary_large_image' }],
		[
			'meta',
			{
				name: 'twitter:image',
				content: 'https://github.com/foomo/gotsrpc/blob/main/docs/public/banner.png?raw=true',
			},
		],
		[
			'meta', { name: 'viewport', content: 'width=device-width, initial-scale=1.0, viewport-fit=cover',
		},
		],
	]
})
