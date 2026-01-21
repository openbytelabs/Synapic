const MODELS = [
    "meta-llama/llama-3.1-405b-instruct:free",
    "nousresearch/hermes-3-llama-3.1-405b:free",
    "qwen/qwen3-235b-a22b:free",
    "openai/gpt-oss-120b:free",
    "meta-llama/llama-3.3-70b-instruct:free",
    "deepseek/deepseek-r1-0528:free",
    "google/gemini-2.0-flash-exp:free",
    "z-ai/glm-4.5-air:free",
    "google/gemma-3-27b-it:free",
    "nvidia/nemotron-3-nano-30b-a3b:free",
    "alibaba/tongyi-deepresearch-30b-a3b:free",
    "allenai/olmo-3.1-32b-think:free",
    "mistralai/mistral-small-3.1-24b-instruct:free",
    "cognitivecomputations/dolphin-mistral-24b-venice-edition:free",
    "openai/gpt-oss-20b:free",
    "mistralai/devstral-2512:free",
    "kwaipilot/kat-coder-pro:free",
    "qwen/qwen3-coder:free",
    "google/gemma-3-12b-it:free",
    "nvidia/nemotron-nano-12b-v2-vl:free",
    "nvidia/nemotron-nano-9b-v2:free",
    "nex-agi/deepseek-v3.1-nex-n1:free",
    "moonshotai/kimi-k2:free",
    "tngtech/deepseek-r1t2-chimera:free",
    "tngtech/deepseek-r1t-chimera:free",
    "tngtech/tng-r1t-chimera:free",
    "mistralai/mistral-7b-instruct:free",
    "qwen/qwen-2.5-vl-7b-instruct:free",
    "amazon/nova-2-lite-v1",
    "xiaomi/mimo-v2-flash:free",
    "google/gemma-3-4b-it:free",
    "google/gemma-3n-e4b-it:free",
    "qwen/qwen3-4b:free",
    "meta-llama/llama-3.2-3b-instruct:free",
    "google/gemma-3n-e2b-it:free",
    "arcee-ai/trinity-mini:free"
];

export const fetchOpenRouterResponse = async (query: string, apiKey: string): Promise<string> => {
    for (const model of MODELS) {
        try {
            const controller = new AbortController();
            const timeoutId = setTimeout(() => controller.abort(), 5000);

            const response = await fetch("https://openrouter.ai/api/v1/chat/completions", {
                method: "POST",
                headers: {
                    "Authorization": `Bearer ${apiKey}`,
                    "Content-Type": "application/json",
                    "HTTP-Referer": "https://synapic.yigitkabak.workers.dev",
                    "X-Title": "Consus"
                },
                body: JSON.stringify({
                    model: model,
                    messages: [
                        { 
                            role: "system", 
                            content: "You are a helpful assistant. You must answer in the same language as the user's query. If the user asks in Turkish, answer in Turkish." 
                        },
                        { 
                            role: "user", 
                            content: query 
                        }
                    ],
                    temperature: 0.7,
                    max_tokens: 1000,
                }),
                signal: controller.signal
            });

            clearTimeout(timeoutId);

            if (response.ok) {
                const data = await response.json();
                if (data && data.choices && data.choices.length > 0) {
                    return data.choices[0].message.content;
                }
            }
        } catch (error) {
            continue;
        }
    }

    return "I'm sorry, all free AI models are currently busy or unable to respond. Please try again shortly.";
};