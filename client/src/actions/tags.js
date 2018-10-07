import { normalize } from "normalizr";
import api from "../api";
import { tagsSchema } from "../schemas";

export const fetchTags = () => () =>
  api.tags.fetchAll().then(tags => normalize(tags, tagsSchema));
export const addTag = tag => () => api.tags.newTag(tag);
