<template>
  <div class="markdown-content" v-html="renderedMarkdown"></div>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import { marked } from 'marked'

interface Props {
  content: string
}

const props = defineProps<Props>()

// 配置marked选项
marked.setOptions({
  breaks: true,
  gfm: true
})

// 使用marked库渲染markdown
const renderedMarkdown = computed(() => {
  if (!props.content) return ''

  try {
    return marked(props.content)
  } catch (error) {
    console.error('Markdown parsing error:', error)
    return props.content
  }
})
</script>

<style lang="less">
@import '@/styles/variables.less';

.markdown-content {
	overflow-wrap: break-word;
	-webkit-font-smoothing: antialiased;
	-moz-osx-font-smoothing: grayscale;

	> * {
		margin: 1rem 0;
	}

	// Title
	h1,
	h2,
	h3,
	h4,
	h5,
	h6 {
		margin: 1rem 0 .5rem;
		padding-bottom: 0.25rem;
		line-height: 1.7;

		&::before {
			display: none;
		}

		& code {
			font-size: inherit !important;
			font-weight: unset;
		}
	}

	h1 {
		margin-bottom: 0.375rem;
		font-size: 1.5rem;
		line-height: 2rem;
	}

	h2 {
		margin-top: 2.25rem;
		padding-bottom: 0.5rem;
		font-size: 1.25rem;
		line-height: 1.75rem;
	}

	h3 {
		margin-top: 1.25rem;
		padding-bottom: 0;
		font-size: 1.125rem;
		line-height: 1.75rem;
	}

	h4 {
		font-size: 1rem;
		line-height: 1.5rem;
	}

	h5 {
		font-size: 0.875rem;
		line-height: 1.25rem;
	}

	h6 {
		margin-top: 0.25rem;
		font-size: 0.75rem;
		line-height: 1rem;
	}

	img {
		margin: 0 auto;
		display: block;

		&[class~='md-image'] {
			overflow: hidden;
			border: 1px @rgb-b3;
			border-radius: 0.375rem;
		}

		&[data-meme-size*='sm'],
		&[data-meme-size*='lg'] {
			display: inline-block;
			margin: 0 .5rem;
		}

		&[data-meme-size*='sm'] {
			width: 36px;
		}

		&[data-meme-size*='lg'] {
			width: 5rem;
		}
	}

	hr {
		border: none @rgb-b3;
		border-top-width: 1px;
		margin: 2.25rem 0;
	}

	code, pre {
		font-family: Fira Code VF, monospace;
		font-weight: 400;
	}

	// 高亮
	code:not(pre code) {
		margin: 0;
		overflow-wrap: break-word;
		border: 1px solid @rgb-b3;
		border-radius: 0.25rem;
		background-color: rgba(@bc, .05);
		padding: .125rem 0.25rem;
		font-size: 0.875rem;
		line-height: 1.25rem;
		box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);

		@supports (color:color-mix(in srgb,transparent,transparent)) {
			background-color: color-mix(in srgb, currentcolor 5%, transparent);
			border-color: color-mix(in srgb, currentcolor 10%, transparent);
		}
	}

	pre {
		position: relative;
		margin: 0.5rem 0;
		height: auto;
		overflow: hidden;
		border-radius: 0.375rem;
		background-color: rgba(var(--bc), 0.06);
		font-size: 0.875rem;
		line-height: 1.25rem;
		transition: height 0.3s ease;

		code {
			margin: 0;
			overflow: auto;
			display: block;
			min-width: 100%;
			width: fit-content;
			padding: 1.25rem 1rem;
			background: transparent;
			border: none;
			border-radius: 0;
			font-size: inherit;
			color: @rgb-bc;
			box-shadow: none;
			white-space: pre;
		}
	}

	a, strong, mark {
		background-image: linear-gradient(to right, rgba(@p, .1) 0%, rgba(@p, .1) 100%);
		background-position: bottom;
		background-repeat: no-repeat;
		background-size: 100% .4rem;
		padding: .125rem;
		border-radius: 0.125rem;
	}

	// Link
	a {
		color: @rgb-p;
		transition: background-size 0.3s ease-out;

		code {
			padding: 0 .25rem !important;
		}

		&:hover {
			background-size: 100% 100%;
		}
	}

	mark {
		background-color: transparent;
		color: @rgb-p;
		font-weight: bold;
		transition: background-size .3s ease-out;

		&:hover {
			background-size: 100% 50%;
		}
	}

	del {
		opacity: 0.4;
	}

	sub, sup {
		font-size: 0.75rem;
		line-height: 1rem;
	}

	video, iframe {
		border-radius: 0.375rem;
		width: 100%;
		height: 500px;
	}

	// Table
	table {
		margin: 0.5rem auto;
		display: inline-block;
		width: auto;
		max-width: 100%;
		border-collapse: collapse;
		overflow-x: auto;
		white-space: nowrap;
		font-size: 0.875rem;
		line-height: 1.25rem;
		color: @rgb-bc;

		thead {
			text-align: left;
			background-color: rgba(var(--bc), 0.06);
		}

		tr:nth-child(2n) {
			background-color: rgba(@bc, .04);
		}

		th,
		td {
			border-width: 1px;
			border-color: @rgb-b3;
			padding: 0.5rem 1rem;
			line-height: 1.25rem;
		}
	}

	// 引用
	blockquote {
		margin: 1rem 0;
		border-left-width: 6px;
		border-color: @rgb-b3;
		padding: 0.5rem 1rem;
		font-size: 0.875rem;
		line-height: 1.25rem;
		font-weight: 500;

		> p {
			margin: .5rem 0;

			&:first-child {
				margin-top: 0;
			}

			&:last-child {
				margin-bottom: 0;
			}
		}
	}

	// Paragraph
	p {
		margin: .75rem 0;
	}

	p, ul, ol {
		line-height: 1.7;
	}

	ul {
		list-style: disc;
	}

	ol {
		list-style: decimal;
	}

	ol,
	ul {
		padding-left: 2rem;

		li {
			margin-bottom: 0;
			list-style-type: inherit;

			&::marker {
				color: @rgb-b3;
			}

			input[type=checkbox] {
				vertical-align: middle;
			}
		}

		ul,
		ol {
			margin: 0.25rem 0;
		}
	}
}
</style>
