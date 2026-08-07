"use client";

import { useQuery } from "@tanstack/react-query";
import { fetchOperators } from "@/lib/api";
import type { OperatorStats } from "@/lib/types";

/** 拉取运营方列表，默认 60s 内缓存 */
export function useOperators() {
  return useQuery<OperatorStats[]>({
    queryKey: ["operators"],
    queryFn: fetchOperators,
    staleTime: 60_000,
    retry: 1,
  });
}
