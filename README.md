# Pre-processink-osquery
TL;DL, pre-processink-osquery is a preprocessing toolkit that helps transform osquery events on a fleet manager, into a machine learning training data set. 

pre-processink-osquery transforms data on a fleet manager into a training set in multiple stages, allowing for automated rules based labeling and manual human labeling. Both kinds of labeling are necessary to cover different use cases and still not make the labeling overwhelming to the administrators.

Here is the pre-processing workflow, it follows:

#### STAGE 0:
 - As of now it queries an ElasticSearch backend of any fleet manager that stores osquery events. (The design being extensible it can be customized to query any data store containing any event source)
 - It queries event of one probe at a time.
 - Labels each event as RED/YELLOW/GREEN based on features, their values and what rules are configured for those values. 
 - If no rules is specified for a particular combination of events or their values the event will not be labeled.
 - All event including labeled and unlabeled ones are written into an output \<probe name>.csv file in textual form.
 - There will be one .csv per probe created in each run.

#### STAGE 1:
 - Merge all the probe specific .csv files into a merged csv file called stage_1.csv. 
 - Data can also be append into a pre-existing stage_1.csv

#### MANUAL LABELING:
 - At this stage human administrator can manually label events that were not labeled or which were incorrectly labeled by the automated rules.

#### STAGE 2:
 - In this stage the labels and features are transformed into numeric values. This will be the training set ready to be fed into a ML toolkit like Tensorflow.
