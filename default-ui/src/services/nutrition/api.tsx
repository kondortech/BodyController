'use server'

import { Api, ApiListIngredientsResponse, HttpResponse, RpcStatus, ApiCreateIngredientResponse, ModelsIngredient, ApiDeleteIngredientResponse, ApiListRecipesResponse, ModelsRecipe, ApiCreateRecipeResponse, ApiDeleteRecipeResponse } from "@/generated/services/nutrition/api";

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


export const createIngredient = async (entity: ModelsIngredient): Promise<ApiCreateIngredientResponse> => {
    const apiClient = new Api({ baseUrl: baseURL });

    const resp: HttpResponse<ApiCreateIngredientResponse, RpcStatus> = await apiClient.nutrition.nutritionCreateIngredient({
        entity: entity
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

export const listRecipes = async (): Promise<ApiListRecipesResponse> => {
    const apiClient = new Api({ baseUrl: baseURL });

    const resp: HttpResponse<ApiListRecipesResponse, RpcStatus> = await apiClient.nutrition.nutritionListRecipes({
        pageSize: 20,
    });

    console.log(resp.error);
    console.log(resp.data);

    return resp.data;
};


export const createRecipe = async (entity: ModelsRecipe): Promise<ApiCreateRecipeResponse> => {
    const apiClient = new Api({ baseUrl: baseURL });

    const resp: HttpResponse<ApiCreateRecipeResponse, RpcStatus> = await apiClient.nutrition.nutritionCreateRecipe({
        entity: entity
    });

    console.log(resp.error);
    console.log(resp.data);

    return resp.data;
};

export const deleteRecipe = async (entityId: string): Promise<ApiDeleteRecipeResponse> => {
    const apiClient = new Api({ baseUrl: baseURL });

    const resp: HttpResponse<ApiDeleteRecipeResponse, RpcStatus> = await apiClient.nutrition.nutritionDeleteRecipe(entityId);

    console.log(resp.error);
    console.log(resp.data);

    return resp.data;
};
