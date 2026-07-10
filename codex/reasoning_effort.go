package codex

import (
	"strings"
)

func normalizeReasoningEffortForModel(args CodexExecArgs) CodexExecArgs {
	args.ModelReasoningEffort = normalizeReasoningEffortValueForModel(args.Model, args.ModelReasoningEffort)
	return args
}

func normalizeReasoningEffortValueForModel(model string, effort string) string {
	trimmedModel := strings.TrimSpace(model)
	trimmedEffort := strings.TrimSpace(effort)
	if trimmedModel == "" || !isMaxOrUltraReasoningEffort(trimmedEffort) {
		return trimmedEffort
	}
	if isGPT56Model(trimmedModel) {
		return trimmedEffort
	}
	return "xhigh"
}

func isMaxOrUltraReasoningEffort(effort string) bool {
	switch effort {
	case "max", "ultra":
		return true
	default:
		return false
	}
}

func isGPT56Model(model string) bool {
	return model == "gpt-5.6" || strings.HasPrefix(model, "gpt-5.6-")
}
