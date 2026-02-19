import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Separator } from "@/components/ui/separator";
import { useMutation, useQuery } from "@tanstack/react-query";
import { getTasks, createTask } from "@/services/tasks";
import { queryClient } from "@/main";
import { useState } from "react";
import Modals from "@/components/Dialog";

const Homepage = () => {
  const [title, setTitle] = useState("");
  const { isPending, isSuccess, isError, error, data } = useQuery({
    queryKey: ["tasks"],
    queryFn: getTasks,
    refetchOnWindowFocus: false,
  });
  const mutation = useMutation({
    mutationFn: createTask,
    onSuccess: () => {
      setTitle("");
      queryClient.invalidateQueries({ queryKey: ["tasks"] });
    },
  });
  return (
    <div className="grid place-items-center min-h-dvh">
      <div className="w-96 border-2 p-4 rounded-sm shadow-sm flex flex-col gap-4">
        <div className="flex gap-2">
          <Input
            placeholder="what is your plan today?"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
          />
          <Button
            onClick={() => {
              mutation.mutate({ title });
            }}
            className="cursor-pointer"
          >
            Add Task
          </Button>
        </div>

        <div>
          <Modals />
        </div>

        {isPending && <p>Loading...🐔</p>}

        {isError && <div className="overflow-auto">error: {error.message}</div>}

        {isSuccess &&
          data.map((task: any) => {
            return (
              <div
                key={task.id}
                className="flex w-full max-w-sm flex-col gap-2 text-sm"
              >
                <dl className="flex items-center justify-between">
                  <dt>{task.title}</dt>
                  <dd className="text-muted-foreground">
                    {`${task.is_completed}`.toUpperCase()}
                  </dd>
                </dl>
                <Separator />
              </div>
            );
          })}
      </div>
    </div>
  );
};

export default Homepage;
