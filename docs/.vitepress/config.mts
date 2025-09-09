import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Go标准库中文文档",
  description: "Go标准库中文文档",
  themeConfig: {
    search: {
      provider: 'local'
    },
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'document', link: '/go-standard-library/' },
    ],

    outline: {
      label: '大纲',
      level: [2, 3],
    },

    sidebar: {
      '/go-standard-library/': [
        {
          text: '常用标准库',
          collapsed: false,
          items: [
            {
              text: 'fmt', link: '/go-standard-library/fmt.md',
            },
            {
              text: 'io', link: '/go-standard-library/io.md',
            },
            { text: 'os', link: '/go-standard-library/os.md' },

            { text: 'strconv', link: '/go-standard-library/strconv.md' },
            { text: 'path', link: '/go-standard-library/path.md' },
            { text: 'runtime', link: '/go-standard-library/runtime.md' },
            { text: 'strings', link: '/go-standard-library/strings.md' },
            { text: 'time', link: '/go-standard-library/time.md' },
          ]
        }
      ],
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/vuejs/vitepress' }
    ]
  }
})
