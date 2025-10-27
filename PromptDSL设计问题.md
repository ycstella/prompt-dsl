1.before是否还需要
如果是文字上的处理，jsonata可能不是特别适用，前处理可能依然需要
比如：存在一个文件放的数据是参考类型的，会通过某一步的结果去找到该文件中的某一项，然后拼接prompt，该项并不会被保存在最终结果中，只是中间的步骤。
2.对于jsonata的使用是引用还是借鉴
3.对于单个任务的for循环嵌套问题，放在单个任务里面好还是外面更好
感觉workflow只管调用会更好。for循环已经是比较具体的业务逻辑了，但是嵌套较深的话after的代码可能不太好写，因为屏蔽了部分实现细节
比如：
after{
        p.afterRet.Id = p.currentData.Id
        p.afterRet.Question = p.currentData.Question
        p.afterRet.Solution = p.currentData.Solution
        // 确保 Steps 长度足够
        if len(p.afterRet.ProcessOut.StepsOut) <= p.it_idx {
            // 补齐
            missing := p.it_idx + 1 - len(p.afterRet.ProcessOut.StepsOut)
            p.afterRet.ProcessOut.StepsOut = append(p.afterRet.ProcessOut.StepsOut, make([]GetpathStepsOut, missing)...)
        }
        log.Println("id:", p.currentData.Process.Steps[p.it_idx].Step, "StepContent:", p.currentData.Process.Steps[p.it_idx].StepContent)
        p.afterRet.ProcessOut.StepsOut[p.it_idx].Step = p.currentData.Process.Steps[p.it_idx].Step
        p.afterRet.ProcessOut.StepsOut[p.it_idx].StepContent = p.currentData.Process.Steps[p.it_idx].StepContent
        p.afterRet.ProcessOut.StepsOut[p.it_idx].Path = p.fixRet.Path
        return nil
}
