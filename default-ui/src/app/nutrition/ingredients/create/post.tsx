'use server'

import { Api, ApiCreateIngredientResponse, HttpResponse, ModelsIngredient, RpcStatus } from "@/generated/services/nutrition/api";

export const createIngredient = async (ingredient: ModelsIngredient): Promise<ApiCreateIngredientResponse> => {
    const apiClient = new Api({ baseUrl: "http://0.0.0.0:8080" });

    const resp: HttpResponse<ApiCreateIngredientResponse, RpcStatus> = await apiClient.nutrition.nutritionCreateIngredient({
        entity: ingredient
    });

    console.log(resp.error);
    console.log(resp.data);

    return resp.data;
};