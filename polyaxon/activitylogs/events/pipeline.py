import activitylogs

from event_manager.events import pipeline

activitylogs.subscribe(pipeline.PipelineCreatedEvent)
activitylogs.subscribe(pipeline.PipelineUpdatedEvent)
activitylogs.subscribe(pipeline.PipelineViewedEvent)
activitylogs.subscribe(pipeline.PipelineArchivedEvent)
activitylogs.subscribe(pipeline.PipelineRestoredEvent)
activitylogs.subscribe(pipeline.PipelineBookmarkedEvent)
activitylogs.subscribe(pipeline.PipelineUnbookmarkedEvent)
activitylogs.subscribe(pipeline.PipelineDeletedTriggeredEvent)
activitylogs.subscribe(pipeline.PipelineCleanedTriggeredEvent)
