import { schema } from "normalizr";

const tag = new schema.Entity("tag");

export const tagsSchema = { tags: [tag] };

export const a = {};
