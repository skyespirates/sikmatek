import { getLimitList } from "@/services/limit";
import { useQuery } from "@tanstack/react-query";
import { LimitTable } from "@/components/LimitTable";

// import { Card, CardContent } from "@/components/ui/card";

const Limit = () => {
  const { isLoading, isError, error, isSuccess, data } = useQuery({
    queryKey: ["limits"],
    queryFn: getLimitList,
  });

  return (
    <div>
      <h1>Daftar Limit</h1>
      {isLoading && <p>Loading...🚀🚀🚀</p>}
      {isError && <p>error: {error.message}</p>}
      {isSuccess && Array.isArray(data) && <LimitTable limits={data} />}
    </div>
  );
};

export default Limit;
