
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  schema: "http://localhost:4000/api",
  documents: ["src/graphql/queries/*.ts"],
  generates: {
    "src/gql/": {
      preset: "client",
      plugins: ["named-operations-object"],
    }
  }
};

export default config;
