# Question:
Reflect on this scenario in the context of the GDPR: What are the potential issues in having the hospital store plaintext private data provided by patients even if they have consented to participate on the experiment and have their data processed? Would these issues be solved by removing the patients' names from their data before storing it?  What are the remaining risks in using Federated Learning with Secure Aggregation as suggested?

## Regarding GDPR:
Briefly looking at the scenario where we store a name and some value about that name, GDPR would only be breached if the data were enough to identify who the person is. Therefore, removing the name would reduce the risk of direct identification, as some value could be anything (and about anyone). However, removing the name alone is not enough, as other data points could still be used to re-identify the individual, especially if combined with external information.

Given that we are dealing with medical data, some value could potentially be a rare disease or something else that could make the person easily identifiable. In the case where we are storing a value from an experiment, it would generally be harder to identify the person in question just from that value alone. However, the risk still depends on the context of the value and whether it could be linked to other identifiable information.

## Remaining risks in using Federated Learning with Secure Aggregation:
- Byzantine Attacks:
Where we have some of the clients sending incorrect data to the server, which in our case degrades the performance of our model.

- Malicious Data:
Adversaries that inject malicious data into the training process of our model to make the model incorrect. This can lead to incorrect predictions, which can be very dangerous when working with something like healthcare.

- Data Leakage:
When sending the shares to the server, it can potentially leak information about the data. Adversaries might then be able to infer sensitive information by analyzing the sent shares, especially if they gain access to lots of shares sent over time.
