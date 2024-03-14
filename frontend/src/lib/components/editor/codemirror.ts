import type { Extension } from '@codemirror/state';
import { basicSetup, EditorView } from 'codemirror';
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
