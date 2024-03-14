import type { MarkdownExtension, BlockContext, Line } from '@lezer/markdown';
import { parseMixed } from '@lezer/common';
import { styleTags, tags } from '@lezer/highlight';
import { foldNodeProp, foldInside, StreamLanguage } from '@codemirror/language';
import { yaml } from '@codemirror/legacy-modes/mode/yaml';

const FENCE_START = '---';
const FENCE_END = '...';

const isFence = (line: Line, fence: string) => {
	return line.text === fence;
};

export const frontmatter: MarkdownExtension = {
	defineNodes: [
		{
			name: 'Frontmatter',
			block: true
		},
		{
			name: 'FrontmatterFence',
			block: false
		}
	],

	props: [
		styleTags({
			Frontmatter: [tags.documentMeta, tags.monospace],

			FrontmatterFence: [tags.processingInstruction]
		}),

		foldNodeProp.add({
			Frontmatter: foldInside
		})
	],

	wrap: parseMixed(({ type, from, to }) => {
		if (type.name !== 'Frontmatter') {
			return null;
		}

		const { parser } = StreamLanguage.define(yaml);

		return {
			parser,
			overlay: [
				{
					from: from + FENCE_START.length,
					to: to - FENCE_END.length
				}
			]
		};
	}),

	parseBlock: [
		{
			name: 'Frontmatter',
			before: 'HorizontalRule',

			parse(ctx: BlockContext, line: Line): boolean | null {
				if (ctx.lineStart === 0 && isFence(line, FENCE_START)) {
					// continue to end fence
					while (ctx.nextLine() && !isFence(line, FENCE_END));

					// check if we hit the end of the file or the end of frontmatter
					if (!isFence(line, FENCE_END)) {
						return false;
					}

					const from = 0;
					const to = ctx.lineStart + FENCE_END.length;

					const fenceStartFrom = from;
					const fenceStartTo = from + FENCE_START.length;

					const fenceEndFrom = to - FENCE_END.length;
					const fenceEndTo = to;

					const fenceStart = ctx.elt('FrontmatterFence', fenceStartFrom, fenceStartTo);
					const fenceEnd = ctx.elt('FrontmatterFence', fenceEndFrom, fenceEndTo);
					const frontmatter = ctx.elt('Frontmatter', from, to, [fenceStart, fenceEnd]);

					ctx.addElement(frontmatter);

					return true;
				}

				return false;
			}
		}
	]
};
