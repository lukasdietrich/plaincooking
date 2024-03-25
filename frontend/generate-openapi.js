import fs from 'node:fs';
import openapiTS from 'openapi-typescript';

const input = '../target/openapi.json';
const output = './src/lib/api/types.gen.ts';

const content = await openapiTS(new URL(input, import.meta.url), {
	transform(schemaObject, metadata) {
		if ('format' in schemaObject && schemaObject.format === 'binary') {
			return 'File';
		}
	}
});

fs.writeFileSync(output, content);
