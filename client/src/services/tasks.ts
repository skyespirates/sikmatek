import { api } from "@/lib/axios";

export async function getTasks() {
  const res = await api.get("/v1/tasks");
  return res.data.data;
}

export async function createTask(task: { title: string }) {
  const res = await api.post("/v1/tasks", task);
  return res.data;
}
