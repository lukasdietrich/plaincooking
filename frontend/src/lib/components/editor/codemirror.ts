import type { Extension } from '@codemirror/state';
import { basicSetup } from 'codemirror';
import { EditorView, keymap } from '@codemirror/view';
import { indentWithTab } from '@codemirror/commands';
import { markdown } from '@codemirror/lang-markdown';
import { frontmatter } from './plugins/frontmatter';

const theme = EditorView.theme({
	'.cm-content': {
		fontFamily: 'inherit',
		fontSize: 'inherit'
	},

	'&.cm-focused': {
		outline: 'none'
	}
});

export function create(element: HTMLElement, content: string, extensions: Extension[] = []) {
	return new EditorView({
		doc: content,
		extensions: [
			basicSetup,
			keymap.of([indentWithTab]),
			theme,
			markdown({
				extensions: frontmatter
			}),
			EditorView.lineWrapping,
			...extensions
		],
		parent: element
	});
}
