'use server'

import { Api, ApiListIngredientsResponse, HttpResponse, RpcStatus, ApiCreateIngredientResponse, ModelsIngredient, ApiDeleteIngredientResponse } from "@/generated/services/nutrition/api";

const baseURL: string = "http://0.0.0.0:8080";

export const listIngredients = async (): Promise<ApiListIngredientsResponse> => {
    const apiClient = new Api({ baseUrl: baseURL });

    const resp: HttpResponse<ApiListIngredientsResponse, RpcStatus> = await apiClient.nutrition.nutritionListIngredients({
        pageSize: 20,
    });

    console.log(resp.error);
    console.log(resp.data);

    return resp.data;
};


export const createIngredient = async (ingredient: ModelsIngredient): Promise<ApiCreateIngredientResponse> => {
    const apiClient = new Api({ baseUrl: baseURL });

    const resp: HttpResponse<ApiCreateIngredientResponse, RpcStatus> = await apiClient.nutrition.nutritionCreateIngredient({
        entity: ingredient
    });

    console.log(resp.error);
    console.log(resp.data);

    return resp.data;
};

export const deleteIngredient = async (entityId: string): Promise<ApiDeleteIngredientResponse> => {
    const apiClient = new Api({ baseUrl: baseURL });

    const resp: HttpResponse<ApiDeleteIngredientResponse, RpcStatus> = await apiClient.nutrition.nutritionDeleteIngredient(entityId);

    console.log(resp.error);
    console.log(resp.data);

    return resp.data;
};

