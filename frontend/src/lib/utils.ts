import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

/**
 * shadcn/ui 标准工具函数：合并 Tailwind CSS 类名
 */
export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}
