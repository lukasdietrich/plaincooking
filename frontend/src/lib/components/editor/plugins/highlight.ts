import type { DecorationSet, EditorView, ViewUpdate } from '@codemirror/view';
import type { SyntaxNodeRef } from '@lezer/common';
import { syntaxTree } from '@codemirror/language';
import { ViewPlugin, Decoration } from '@codemirror/view';
import { RangeSetBuilder } from '@codemirror/state';

interface DecorationRule {
	class: string;
	line?: boolean;
}

const decorateNodes = (builder: RangeSetBuilder<Decoration>) => {
	const rules: Record<string, DecorationRule> = {
		ATXHeading1: {
			class: 'cm-heading-1'
		},

		HorizontalRule: {
			class: 'cm-hr',
			line: true
		},

		InlineCode: {
			class: 'cm-ingredient'
		},

		CodeMark: {
			class: 'cm-codemark'
		}
	};

	return ({ from, to, name }: SyntaxNodeRef) => {
		const rule = rules[name];
		if (rule) {
			if (rule.line) {
				builder.add(from, from, Decoration.line({ class: rule.class }));
			} else {
				builder.add(from, to, Decoration.mark({ class: rule.class }));
			}
		}
	};
};

class HighlightPlugin {
	decorations: DecorationSet = Decoration.none;

	constructor(view: EditorView) {
		this.updateDecorations(view);
	}

	update({ docChanged, viewportChanged, startState, state, view }: ViewUpdate) {
		if (docChanged || viewportChanged || syntaxTree(startState) != syntaxTree(state)) {
			this.updateDecorations(view);
		}
	}

	private updateDecorations({ state, visibleRanges }: EditorView) {
		const builder = new RangeSetBuilder<Decoration>();

		for (const { from, to } of visibleRanges) {
			syntaxTree(state).iterate({
				from,
				to,
				enter: decorateNodes(builder)
			});
		}

		this.decorations = builder.finish();
	}
}

export const highlight = ViewPlugin.fromClass(HighlightPlugin, {
	decorations(plugin: HighlightPlugin) {
		return plugin.decorations;
	}
});
