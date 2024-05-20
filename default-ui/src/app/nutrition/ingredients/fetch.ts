'use server'

import { Api, ApiListIngredientsResponse, HttpResponse, RpcStatus } from "@/generated/services/nutrition/api";

export const fetchIngredients = async (): Promise<ApiListIngredientsResponse> => {
    const apiClient = new Api({
        baseUrl: "http://0.0.0.0:8080",
    });

    const resp: HttpResponse<ApiListIngredientsResponse, RpcStatus> = await apiClient.nutrition.nutritionListIngredients({
        pageSize: 20,
    });

    console.log(resp.error);
    console.log(resp.data);

    return resp.data;
};