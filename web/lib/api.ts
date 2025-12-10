const API_BASE_URL =
  process.env.NEXT_PUBLIC_API_BASE_URL || "http://127.0.1.1:8080";

export type Site = {
  id: number;
  nome: string;
  url: string;
  status: number;
};

export const API_ROUTES = {
  SITES: `${API_BASE_URL}/`,
};
