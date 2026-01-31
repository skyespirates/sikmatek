import { useState, type FormEvent } from "react";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useMutation } from "@tanstack/react-query";
import { Login as LoginUser } from "@/services/auth";
import { useNavigate } from "react-router";
const Login = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const mutation = useMutation({
    mutationFn: LoginUser,
    onSuccess: (data) => {
      localStorage.setItem("token", data.token);
      navigate("/");
    },
    onError: ({ response }: { response: { data: string } }) => {
      console.log(response.data);
    },
  });

  function handleSubmit(e: FormEvent<HTMLFormElement>) {
    e.preventDefault();

    const data = {
      email,
      password,
    };
    mutation.mutate(data);
  }
  return (
    <div className="h-dvh  grid place-items-center">
      <div className="p-8 w-96 border-2 rounded-sm shadow-md flex flex-col items-center gap-6">
        <h1 className="text-3xl">Login</h1>
        <form onSubmit={handleSubmit} className="flex flex-col gap-2 w-full">
          <div>
            <Input
              className=""
              placeholder="someone@email.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
          </div>
          <div>
            <Input
              type="password"
              placeholder="Enter you password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
          </div>
          <div className="flex">
            <Button type="submit" className="flex-1 cursor-pointer">
              Login
            </Button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Login;
