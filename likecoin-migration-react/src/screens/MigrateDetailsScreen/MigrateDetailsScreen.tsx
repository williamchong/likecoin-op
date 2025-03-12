import { useQuery } from "@tanstack/react-query";
import { useParams } from "react-router";

import useGetMigrationRecord, {
  queryKey,
} from "../../apis/useGetMigrationRecord";

export default function MigrateDetailScreen() {
  const { cosmosTxHash } = useParams();

  const getMigrationRecord = useGetMigrationRecord(cosmosTxHash!);

  const { data, isLoading } = useQuery({
    queryKey: queryKey(cosmosTxHash!),
    queryFn: () => {
      return getMigrationRecord();
    },
    refetchInterval(query) {
      if (query.state.error != null) {
        return undefined;
      }
      if (
        query.state.data == null ||
        query.state.data.status == "success" ||
        query.state.data.status == "failed"
      ) {
        return undefined;
      }
      return 2000;
    },
  });

  const shouldShowLoading = isLoading || data?.status === "pending";

  return (
    <div className="grid grid-rows-[20px_1fr_20px] items-center justify-items-center min-h-screen p-8 pb-20 gap-16 sm:p-20 font-[family-name:var(--font-geist-sans)]">
      <main className="flex flex-col gap-8 row-start-2 items-center sm:items-start">
        <p>Cosmos Tx Hash: {data?.migration_record.cosmos_tx_hash}</p>
        <p>Cosmos Address: {data?.migration_record.cosmos_address}</p>
        <p>Ethereum Address: {data?.migration_record.eth_address}</p>
        <p>Ethereum Tx Hash: {data?.migration_record.eth_tx_hash}</p>
        <p>Status: {data?.status}</p>
      </main>
      {shouldShowLoading ? (
        <div className="fixed top-0 left-0 w-full h-full bg-white/90 flex items-center justify-center">
          Loading
        </div>
      ) : null}
    </div>
  );
}
