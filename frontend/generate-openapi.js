import fs from 'node:fs';
import ts from 'typescript';
import openapiTS, { astToString } from 'openapi-typescript';

const input = '../target/openapi.json';
const output = './src/lib/api/types.gen.ts';

const ast = await openapiTS(new URL(input, import.meta.url), {
	transform(schemaObject) {
		if ('format' in schemaObject && schemaObject.format === 'binary') {
			return ts.factory.createTypeReferenceNode(ts.factory.createIdentifier('Blob'));
		}
	}
});

const content = astToString(ast);

fs.writeFileSync(output, content);
