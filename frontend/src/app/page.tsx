"use client";

import { useState } from "react";
import { RequestForm } from "@/components/request/RequestForm";
import { QuoteSummary } from "@/components/request/QuoteSummary";
import { OperatorPicker } from "@/components/operator/OperatorPicker";
import { ErrorBanner } from "@/components/ui/ErrorBanner";
import { useQuote, useCreatePlan, DEFAULT_DEADLINE_MINUTES } from "@/hooks/useQuote";
import { useOperators } from "@/hooks/useOperators";
import { HOME, COPY } from "@/lib/copy";
import type { OperatorStats } from "@/lib/types";

/** 首页: 选择运营方 → 输入请求 → 获取报价 → 生成担保计划 (文档 §6.1) */
export default function HomePage() {
  const [amountWei, setAmountWei] = useState<string | undefined>(undefined);
  const [selectedOperator, setSelectedOperator] = useState<OperatorStats | undefined>(undefined);

  const operatorsQuery = useOperators();
  const quoteQuery = useQuote(amountWei, !!amountWei);
  const createPlan = useCreatePlan();

  const quote = quoteQuery.data;

  // Auto-select first operator when list loads
  if (!selectedOperator && operatorsQuery.data?.length) {
    setSelectedOperator(operatorsQuery.data[0]);
  }

  const handleCreate = () => {
    if (!quote) return;
    createPlan.mutate({
      inputAmount: quote.inputAmount,
      expectedOutput: quote.expectedOutput,
      deadlineMinutes: DEFAULT_DEADLINE_MINUTES,
    });
  };

  // Merge selected operator info into quote for display
  const quoteWithOperator = quote && selectedOperator
    ? { ...quote, operator: selectedOperator }
    : quote;

  return (
    <div className="space-y-6">
      {/* 标题区 */}
      <div className="pt-4">
        <h1 className="bg-gradient-to-r from-primary to-accent-tech bg-clip-text text-2xl font-bold text-transparent sm:text-[28px]">
          {HOME.title}
        </h1>
        <p className="mt-2 text-sm text-text-secondary">{COPY.tagline}</p>
      </div>

      {/* 运营方选择器 */}
      <OperatorPicker
        operators={operatorsQuery.data ?? []}
        selected={selectedOperator?.address}
        onSelect={setSelectedOperator}
        loading={operatorsQuery.isLoading}
      />

      {/* 请求表单 */}
      <RequestForm onQuote={setAmountWei} loading={quoteQuery.isFetching} />

      {/* 报价预览 */}
      {quoteWithOperator && (
        <QuoteSummary
          quote={quoteWithOperator}
          onCreatePlan={handleCreate}
          creating={createPlan.isPending}
        />
      )}

      {/* 错误区 */}
      <ErrorBanner
        error={quoteQuery.isError ? quoteQuery.error : createPlan.isError ? createPlan.error : null}
        onRetry={quoteQuery.isError ? () => quoteQuery.refetch() : undefined}
      />
    </div>
  );
}
