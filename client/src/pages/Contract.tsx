import { ContractTable } from "@/components/ContractTable";
import { getDaftarKontrak } from "@/services/contract";
import { useQuery } from "@tanstack/react-query";

const Contract = () => {
  const { isLoading, isError, error, isSuccess, data } = useQuery({
    queryKey: ["kontrak"],
    queryFn: getDaftarKontrak,
  });

  return (
    <div>
      <p>Daftar Kontrak</p>
      {isLoading && <p>Loading... 🚀🚀🚀</p>}
      {isError && <p>error: {error.message}</p>}
      {isSuccess && <ContractTable contracts={data} />}
    </div>
  );
};

export default Contract;
