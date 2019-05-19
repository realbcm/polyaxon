import tracker

from event_manager.events import pipeline_run

tracker.subscribe(pipeline_run.PipelineRunCreatedEvent)
tracker.subscribe(pipeline_run.PipelineRunUpdatedEvent)
tracker.subscribe(pipeline_run.PipelineRunCleanedTriggeredEvent)
tracker.subscribe(pipeline_run.PipelineRunViewedEvent)
tracker.subscribe(pipeline_run.PipelineRunArchivedEvent)
tracker.subscribe(pipeline_run.PipelineRunRestoredEvent)
tracker.subscribe(pipeline_run.PipelineRunDeletedEvent)
tracker.subscribe(pipeline_run.PipelineRunDeletedTriggeredEvent)
tracker.subscribe(pipeline_run.PipelineRunStoppedEvent)
tracker.subscribe(pipeline_run.PipelineRunStoppedTriggeredEvent)
tracker.subscribe(pipeline_run.PipelineRunResumedEvent)
tracker.subscribe(pipeline_run.PipelineRunResumedTriggeredEvent)
tracker.subscribe(pipeline_run.PipelineRunRestartedEvent)
tracker.subscribe(pipeline_run.PipelineRunRestartedTriggeredEvent)
tracker.subscribe(pipeline_run.PipelineRunSkippedEvent)
tracker.subscribe(pipeline_run.PipelineRunSkippedTriggeredEvent)
tracker.subscribe(pipeline_run.PipelineRunNewStatusEvent)
tracker.subscribe(pipeline_run.PipelineRunSucceededEvent)
tracker.subscribe(pipeline_run.PipelineRunFailedEvent)
tracker.subscribe(pipeline_run.PipelineRunDoneEvent)
